package user

import (

	"appengine"	
	"appengine/datastore"	
	//"net/http"
	"time"
	"fmt"
)


type User struct {

	Id 	    string
	Name    string
	Link    string
	Picture string
	Gender  string
	
	Created time.Time
}




func Create(user *User, c appengine.Context) (*User){
		
	existed, err := Get(user.Id, c)
	if err != nil {
	
		c.Warningf(err.Error())
			
		user.Created = time.Now()
		
		//k := datastore.NewKey(c, "User", user.Id, 0, nil)
	    k := datastore.NewIncompleteKey(c, "User", nil)       
	    
	    _, err := datastore.Put(c, k, user)
	    if err != nil {
	    	c.Errorf(err.Error())
	    }
	    
	    return user
        
    }
    
    return existed 
      
    	
    
    
}


func Get(id string, c appengine.Context) (user *User, err error) {

	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	
	q := datastore.NewQuery("User").Filter("Id =", id).Order("-Created").Limit(1)
    
	var users []User
	    
	if _, err := q.GetAll(c, &users); err != nil {	
				
		c.Errorf(err.Error())
		return nil, err
	}
	
	
	if len(users) == 0 {
	
		return nil, fmt.Errorf("User not found: %s", id)				
	}
		
	return &users[0], nil
	
}
