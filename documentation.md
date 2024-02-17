This project is written in Golang
I have used GIN web framework, GORM orm and MySql database

I have added the shchema in deployments/deoploytable.sql file

I have followed design patterns to make the code more extensible, flexible and easy to understand. 

I have followed MVC design pattern for this project.

The user can send request via two apis endpoints
    1. curl --location 'localhost:8080/api/temperature' \
    --header 'Content-Type: application/json' \
    --data '{
        "sensor_id":556,
        "temperature_value":40,
        "time_stamp": 1644931200
    }'

    This endpoint handles new add temperature request coming from multipl sensor points. As a request body sensor needs to send the sensor_id, temperature_value and time_stamp. sensor_id and temeperature_value is mandatory but time stamp is not. If timestamp is not sent I am adding current timestamp.

    2. curl --location --request GET 'localhost:8080/api/aggregate' \
    --header 'Content-Type: application/json' \
    --data '{
        "sensor_id":556,
        "start_time":"01:04:05",
        "end_time":"23:22:00",
        "start_date":"2024-02-15",
        "end_date":"2024-02-17"
    }'

    This endpoint fetches the data of a particular sensor, optionally given date or time range. The sensor_id is a mandatory field. if the date and time are not given then this endpoint fetches data of the sensor. If only date is not given and time given then it returns entries of todays data for that sensor within that time range.

The apis now calls the controller functions.

In the controller folder I have added temperature controller which have methods to add new temperature and get the aggregated temeperature. controller operate on the data transfer objects to communicate with the client. 

Controller calls the service layer functions. Service layer functions holds the business logic of the project and for interacting with the database service layer functions call the repository layer functions.

In the repository layer functions I have added ORM methods to communicate with database. 

I have added a cron job that runs hourly and fetches data from the temperature table of the temperature of the last one hour. and does the aggregation on them and stored the data of every sensor. this is done by robfig/cron package. 

I have added interfaces for all the layers to communicate with other layers so the layers are not directly dependant on each other. 

I have declared three tables. 
1. Sensor table: - this table holds the data of the sensors. For now it is just holding the id of the sensor.

2. Temperature table - this table hold the temepratures of sensor points, timestamps and the sensor id (which is a foreign key of the sensor table)

3. Aggregated_temperature table - this table holds the aggregated value for temperatures. data on this table is filled by the cron job. This cron job runs once in every hour. And fetches the temperatures of last one hour and does aggregation on them such as max, min, avg. 

