# keycloak-go-server

####################

=> Run into terminal "go mod tidy" // install all modules

=> At ./ of project run docker with docker compose // pull keycloak docker img and run at http://localhost:8086

=> Open navigator at http://localhost:8086 (username: admin, password: admin)

=> Configur new realm name: 'medium'

=> Client 
    => create new client name: "my-go-service"

=> Client 
    => my-go-service 
        => settings 
            => {
                "Enabled": true, 
                "Acces type": confidential, 
                "Direct Acces Garants Enabled": true, 
                "Service Account Enabled": true, 
                "Backchannel Logout Session Required": true
                }

=> Client 
    => my-go-service 
        => credentials 
            => copy secret and add into ./client/keycloak.go/clientSecret 

=> Client 
    => my-go-service 
        => service account roles 
            => client roles 
                => select realm managment and add view users

=> Users 
    => create new user whith password and more users if you want

=> For run server "go run ." // run project at http://localhost:8081

########################

=> Request whith postman or insomnia 

    => POST - http://localhost:8081/login 
        json {
	        "username": "myuser",
	        "password":"xxx"
        }
        // this login request get acces token (save this token you need for other get request)

    => GET - http://localhost:8081/users
        => auth : Bearer => token : your acces token
        // get all realm users
        
    => GET - http://localhost:8081/user
        => auth : Bearer => token : your acces token
        // get current user infos 
        
        
