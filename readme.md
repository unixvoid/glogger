glogger
-------
- glogger is a logging framework written in golang.  
- usage:  
  - to impliment glogger, first add the import: `import "git.unixvoid.com/mfaltys/glogger"`
  - then initialize the logger:  
    ```
	if *logLevel {  
		glogger.LogInit(os.Stdout, os.Stdout, os.Stderr)  
	} else {  
		glogger.LogInit(os.Stdout, ioutil.Discard, os.Stderr)  
	}
	```  
  - to then use the logger, write your print statements like the following:  
	`glogger.Info.Println("Produciton app running on 8080")`  
    `glogger.Debug.Printf("DEBUG :: serving response\n")`  
    `glogger.Error.Println("DEBUG :: serving response")`  
