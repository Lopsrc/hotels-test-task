# hotels-test-task

## About it
This is test task for Hotels.

### Technologies

>Golang version 1.22.2
>
>Docker version 26.1.0

## Using

### Clone the repositiry
```
git clone https://github.com/Lopsrc/hotels-test-task
```

### Running

Run on local machine:
```
go run main.go
```

Run in docker:
```
# Building docker container.
docker build -t hotels .

# Running docker container.
docker run -it hotels
```