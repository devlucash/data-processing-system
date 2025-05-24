# Data Processing System 

This is a basic CRUD web application written in Go that simulates data generation and processing.

## 📁 Project Structure using MVC design
```bash
data-processing-system/
├── controller/   # Handles HTTP routes and logic
│ └── home.go
├── model/        # Generates data objects for processing
│ └── solar.go
├── view/         # HTML templates
│ └── index.html
├── main.go       # Entry point of the application
├── go.mod        # Go module file
└── README.md
```
## Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/devlucash/data-processing-system.git
cd data-processing-system
```

### 2. run the app
```
go run main.go
```
or using the launch.json with VScode built in run

### 3. View in browser
Visit: ``` http://localhost:8080 ```

## Project features
- Simulate solar/battery/grid data switching logic
- Add NoSQL data storage
- Implement full CRUD (Create, Read, Update, Delete) functionality
- REST API endpoints
- Authentication and session handling

## Technology 
- Go
- NoSQL (MongoDB?)
- UI framwork (maybe Bootstrap to keep it simple)




