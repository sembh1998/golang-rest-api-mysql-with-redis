# golang-rest-api-mysql-with-redis
 
 With this i want to show my logic view of how 
 to work with redis and mantain it updated always
 with no risk of getting outdated data

 Here it is an image of what can be deployed with this repository as an guide

 ![image](https://github.com/sembh1998/golang-rest-api-mysql-with-redis/blob/main/readme.images/docker-compose-to-show-my-logic.png) 

My logic of how to work with redis is this:

> Create a standard to save the last date of when a table had have a change and save it the now date in unixNano format, example:
  SET "mysql.table.last_change_date" 1233455678234

> Create a standard to save the date of when the sql consult was do, example:
  GET "mysql.getallproducts.consult_date" 12381952498174

A consult can use one or many tables, so a function is need with this structure

func validateCacheData(tables []string, consult string){
    //consult last_change_date of every table implicated in the consult
    // compare the dates, if the date of the consult is grader than the dates of the tables return from redis directly
    // if the date of any of the tables is grader than the date of the consult then made the consult again and save it again and 
    // update the date of the consult
  }

example use:

consult := "select * from users u left join mails m on u.id = m.user_id"
tables := []string{"users","mails"}

validateCacheData(tables, consult)

