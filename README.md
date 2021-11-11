This is a golang program that stores data in a MongDB for possible future research.

This project uses the golang mongodb-driver 
To get the driver enter the following command at the command prompt in the project directory

go get go.mongodb.org/mongo-driver/mongo

One package was created for this project: github.com/afbaum/mongoGoStudy

Several functions or methods were used and at least one returns a value:
1. MongoConnection is a function the creates a connection to the mongo database, it returns a value which is a pointer to the mongo client
2. dataEntry function collects data from the html form, it connects to the database and inserts the data into the database
3. handleRequests methonds handles the functions requested based on the URL entered.

I added a loop which asks for a user name of greater than five characters.  If you do not enter enough characters the program will ask again.  You may also type 'q' to quit.  This componenet was only added for the purposes of the Code Lou project requirements.  
