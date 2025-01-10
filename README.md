# coin-App


# To Run This

step 1 :

create a database in postgres.

step 2 :

Add the database name and your postgres password to the dsn string inside the ConnectDB function from utils/db/database.go

# port running on :8000

# Api calls and Request Models Below

# Health check
  http://localhost:8000/ping   
  {
# 1: request model of Add Coin 
  http://localhost:8000/updatecoin
  
{

   "name":"Etherium",
   "image":"https://example.com/bitcoin.png"
   "expiry_date":"<your_expiration_time>"  default expiration time is one day

}


# 2: request model of Add View by name or Id 
  http://localhost:8000/viewbyname/or/id

{
  
   "name":"Binance" 
  
}  

# or  

{

  "id": 1    

}

# 3: request model of List All Non Expired Coins 
 http://localhost:8000/view/all

empty

# 4: request model of Listing Expired Coin Logs
 http://localhost:8000/view/expiredcoins/log

empty
  }



