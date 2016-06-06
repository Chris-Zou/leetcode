package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	TarFile := "test.tar.gz"
	src := "test.txt"
	//dstDir := "test_ext"

	if err := Tar(src, TarFile, false); err != nil {
		fmt.Println(err)
	}

	/*if err := UnTar(TarFile, dstDir); err != nil {
		fmt.Println(err)
	}*/
}

func Tar(src string, dstTar string, failIfExist bool) (err error) {
	src = path.Clean(src)

	if !Exists(src) {
		return errors.New("directory does not exist:" + src)
	}

	if FileExists(dstTar) {
		if failIfExist {
			return errors.New("target file exist:" + dstTar)
		} else { // rewrite the file that exist
			if err := os.Remove(dstTar); err != nil {
				return err
			}
		}
	}
	// create new target file
	fw, er := os.Create(dstTar)
	if er != nil {
		return err
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)

	tw := tar.NewWriter(gw)
	defer func() {
		if er := tw.Flush(); er != nil {
			fmt.Println("flush file failed")
		}
		if er := tw.Close(); er != nil {
			fmt.Println("close file failed")
			err = er
		}
	}()

	fi, er := os.Stat(src)
	if er != nil {
		return er
	}

	srcBase, srcRelative := path.Split(path.Clean(src))

	if fi.IsDir() {
		tarDir(srcBase, srcRelative, tw, fi)
	} else {
		fmt.Println("tar file")
		tarFile(srcBase, srcRelative, tw, fi)
	}

	return nil
}

func tarDir(srcBase, srcRelative string, tw *tar.Writer, fi os.FileInfo) (err error) {
	srcFull := srcBase + srcRelative

	last := len(srcRelative) - 1
	if srcRelative[last] != os.PathSeparator {
		srcRelative += string(os.PathSeparator)
	}

	fis, err := ioutil.ReadDir(srcFull)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		if fi.IsDir() {
			tarDir(srcBase, srcRelative+fi.Name(), tw, fi)
		} else {
			tarFile(srcBase, srcRelative+fi.Name(), tw, fi)
		}
	}
	if len(srcRelative) > 0 {
		hdr, er := tar.FileInfoHeader(fi, "")
		if er != nil {
			return er
		}
		hdr.Name = srcRelative

		if er = tw.WriteHeader(hdr); er != nil {
			return er
		}
	}

	return nil
}

// 因为要在 defer 中关闭文件，所以要单独创建一个函数
func tarFile(srcBase, srcRelative string, tw *tar.Writer, fi os.FileInfo) (err error) {
	// 获取完整路径
	srcFull := srcBase + srcRelative

	// 写入文件信息
	hdr, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		fmt.Println("line 124:", err)
		return err
	}
	hdr.Name = srcRelative

	if err = tw.WriteHeader(hdr); err != nil {
		fmt.Println("line 130:", err)
		return err
	}

	// 打开要打包的文件，准备读取
	fr, err := os.Open(srcFull)
	if err != nil {
		fmt.Println("line 137:", err)
		return err
	}
	defer fr.Close()

	// 将文件数据写入 tw 中
	if _, err = io.Copy(tw, fr); err != nil {
		fmt.Println("line 144:", err)
		return err
	}
	return nil
}

func UnTar(srcTar string, dstDir string) (err error) {
	// 清理路径字符串
	dstDir = path.Clean(dstDir) + string(os.PathSeparator)

	// 打开要解包的文件
	fr, err := os.Open(srcTar)
	if err != nil {
		return err
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)

	// 创建 tar.Reader，准备执行解包操作
	tr := tar.NewReader(gr)

	// 遍历包中的文件
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			return err
		}

		// 获取文件信息
		fi := hdr.FileInfo()

		// 获取绝对路径
		dstFullPath := dstDir + hdr.Name

		if hdr.Typeflag == tar.TypeDir {
			// 创建目录
			os.MkdirAll(dstFullPath, fi.Mode().Perm())
			// 设置目录权限
			os.Chmod(dstFullPath, fi.Mode().Perm())
		} else {
			// 创建文件所在的目录
			os.MkdirAll(path.Dir(dstFullPath), os.ModePerm)
			// 将 tr 中的数据写入文件中
			if er := unTarFile(dstFullPath, tr); er != nil {
				return er
			}
			// 设置文件权限
			os.Chmod(dstFullPath, fi.Mode().Perm())
		}
	}
	return nil
}

// 因为要在 defer 中关闭文件，所以要单独创建一个函数
func unTarFile(dstFile string, tr *tar.Reader) error {
	// 创建空文件，准备写入解包后的数据
	fw, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer fw.Close()

	// 写入解包后的数据
	_, err = io.Copy(fw, tr)
	if err != nil {
		return err
	}

	return nil
}

// 判断档案是否存在
func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

// 判断文件是否存在
func FileExists(filename string) bool {
	fi, err := os.Stat(filename)
	return (err == nil || os.IsExist(err)) && !fi.IsDir()
}

// 判断目录是否存在
func DirExists(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}
