Title:
8gaming

Description:
Generate Event data and write to file in batches.
To Run in terminal go run 8gaming --number-of-groups=100 --batch-size=20 --interval=3 --output-directory=your directory

Logic:
Starting in main.go i get the inputs and parse them for use, in the handleInputs method i do the batch calculation and calculate the percentages for every type needed for 
the batch, i then round those figures to handle decimals. There is a dateTime variable which is the time data was generated, i use this as well as the rounded percentage values
to call a method named createViewedData in structhelpers class this class gererates the data and send it to a function called writeAsBatch in filehelpers class which creates a 
file based on the data and directory.

ErrorHandling:
In terms of ErrorHandling for inputs the cli will return if value is incorrect.
There is a check for valid directory and will exit status 2 with a error message.
All other failures will exit and log the error.

HardCoded Values:
All Hardcoded values would be configs in Aws parameter store.

Testing:
I have very little testing framework experience.
As discussed in the interview we only test our code to see if it works and then it gets sent to business users to test scenarios.