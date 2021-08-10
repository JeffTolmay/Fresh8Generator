Title:
Fresh8Poller

Description:
Monitor a location for new files and print the various events.
To Run in terminal go run fresh8Poller

Logic:
Starting in main.go i start the watcher which checks the directory for new files in a certain directory.
If a file is created and has the event and .json i try to read the file line by line using the readfile method.
If the file is correct i will unmarshal it to json and loop over the values to get the totals.

ErrorHandling:
All failures will exit and log the error.

HardCoded Values:
All Hardcoded values would be configs in Aws parameter store.

Testing:
I have very little testing framework experience.
As discussed in the interview we only test our code to see if it works and then it gets sent to business users to test scenarios.