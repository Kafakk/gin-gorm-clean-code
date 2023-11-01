# gin-gorm-clean-code

 Create Database table
```
CREATE TABLE students (
	id serial PRIMARY KEY,
	first_name VARCHAR ( 50 ) UNIQUE NOT NULL,
	last_name VARCHAR ( 50 ) NOT NULL,
	age int NOT NULL,
	sex VARCHAR ( 255 ) UNIQUE NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL
);
```

