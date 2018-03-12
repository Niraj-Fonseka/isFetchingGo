Experiment to reduce thirdparty API calls using Go ( Channels , Routines )

API - 
    This api acts as an api hub to fetch data from other third party sources

APP - 
    This application will be open to a multiple users. And the APP will call the API 

Assumptions - 
    
    - Every third party api call is getting real time data ( time sensitive ) and Rate Limited ( 1000 calls per day )
        ie :
            GET /population - Returns live population in the world. Changes every minute. 
            
    - Every third party api call is costly


Problem - 
   
    - If multiple people tend to hit the same API endpoint ( GET /population ) at the same time based on a some external trend ( ie : BRAKING NEWS - Scientists suggests that population growth have slowed down 10 times compared to last year). We will be making many api calls to get about the same data.
        ie : at 10:20:00 AM 50 people ask for what the population is 
                Current result : make 50 calls to GET /population 


Expected Solution -
    
    - if the request was done for the same time stamp. Do just one call and use the result to fullfill all the other 49 requests. 


Findings - 
    ... TBD
    