## This is a golang program that stores data in a MongDB for possible future research.

This project uses the golang mongodb-driver 
To get the driver enter the following command at the command prompt in the project directory

**go get go.mongodb.org/mongo-driver/mongo**

One package was created for this project: github.com/afbaum/mongoGoStudy

 #### Three different pages were created for this project
 1. localhost:8080/ 
    this it the static front page to the project
2. localhost:8080/form
    this page has a html form.  Data entered into this form is stored in the MongodB database.
3. localhost:8080/infoPage
    the database in queried for only the male subjects and the IOIHA results are printed onto the page.

#### Several functions or methods were used and at least one returns a value:
1. MongoConnection is a function the creates a connection to the mongo database, it returns a value which is a pointer to the mongo client
2. dataEntry function collects data from the html form, it connects to the database and inserts the data into the database
3. handleRequests methonds handles the functions requested based on the URL entered.

#### Extra Features
1. In the pipeline function I connected to an external database, read information from the database in BSON format.  I then looped through the map of data to obtain only the results for the ioiha measure.  I converted the ioiha data into a string data type and populated a slice with the results.  This slice is then used by the a html template in which a list of the results is presented in the browser.

There obviously needs to be a better appearnace to my html pages, however I think the most important next step would be to clean up my main.go file and separate my functions out into different files.  By different components into different files it would be much easier to go back through and clean up the code.
