# keycloak-go-server

####################

=> go mod tidy // install all modules

=> at ./ of project run docker with docker compose // pull keycloak docker img and run at http://localhost:8086

=> open navigator at http://localhost:8086 (username: admin, password: admin)

=> configur new realm name: 'medium'

=> client => create new client name: "my-go-service"

=> client => my-go-service => settings => "Enabled": true, "Acces type": confidential, "Direct Acces Garants Enabled": true, "Service Account Enabled":        true, "Backchannel Logout Session Required": true

=> client => my-go-service => credentials => copy secret and add into ./client/keycloak.go => clientSecret 

=> Users => create new user whith password 

=> go run *.go // run project at http://localhost:8081

########################

=> request whith postman or insomnia 
    => POST - http://localhost:8081/login 
    
        json {
	        "username": "myuser",
	        "password":"xxx"
        }
        // this login request get acces token (save this token you need for other get request)

    => GET - http://localhost:8081/docs
        => auth : Bearer => token : your acces token
        // get docs with middleware for auth exercice
        
    => GET - http://localhost:8081/user
        => auth : Bearer => token : your acces token
        // get current user infos 
        
        
