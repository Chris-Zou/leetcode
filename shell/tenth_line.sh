# Read from the file file.txt and output the tenth line to stdout.
#awk 'NR == 10' file.txt #using awk, accepted
#sed -n '10p' file.txt   #using sed, accepted
#tail -n+10 file.txt | head -n1 #using head and tail, accepted
head -n10 file2.txt | tail -n+10 #using head and tail, accepted
