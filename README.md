This is a golang program that stores data for possible research in a MongoDB.

One package was created for this project: github.com/afbaum/mongoGoStudy

Several functions or methods were used and at least one returns a value:
1. MongoConnection is a function the creates a connection to the mongo database, it returns a value which is a pointer to the mongo client
2. dataEntry function collects data from the html form, it connects to the database and inserts the data into the database
3. handleRequests methonds handles the functions requested based on the URL entered.

Three items from the Features List
1. This program reads data from a Mongo Database and uses that data by reporting it in the frontend.
2. Data from the MongoDB is visualized in a graph
3. Data is analyzed and information is displayed about the data 
4. the golang log is implimented and used throughout the code to identify errors