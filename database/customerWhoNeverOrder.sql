#SELECT Name as Customers FROM Customers c WHERE c.Id NOT IN (SELECT CustomerId FROM Orders o);
SELECT Name as Customers FROM Customers c WHERE NOT EXISTS (SELECT CustomerId FROM Orders o WHERE o.CustomerId = c.id);
