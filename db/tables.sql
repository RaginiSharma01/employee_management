CREATE TABLE employees_data (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    department VARCHAR(100),
    salary NUMERIC(10,2),
    joining_date TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- selecting all employees 
SELECT *FROM employees_data;


--employee by salary 
SELECT *FROM employees_data WHERE salary > 50000;


-- by department 
SELECT *FROM employees_data WHERE department = "marketing" 


-- counting
SELECT department, COUNT(*) AS employee_count FROM employees_data GROUP by department;

--by joining date
SELECT *FROM employees_data WHERE joining_date>=NOW() - INTERVAL'1-day'

--top 5 salary
SELECT *FROM employees_data ORDER By salary DESC LIMIT 5

-- random 10 user adding 

CREATE OR REPLACE FUNCTION add_random_employees()
RETURNS void
LANGUAGE plpgsql
AS $$
DECLARE
    names TEXT[] := ARRAY[
       'nameesh','john','michael','sarah','emily','david','jessica','daniel','laura','robert'
    ];

    departments TEXT[] := ARRAY[
        'Engineering','Finance','HR','Marketing','Sales'
    ];

    i INT;
    random_name TEXT;
    random_dept TEXT;
    random_salary INT;
BEGIN
    FOR i IN 1..10 LOOP

        random_name := names[floor(random()*array_length(names,1) + 1)];
        random_dept := departments[floor(random()*array_length(departments,1) + 1)];
        random_salary := floor(random()*70000 + 30000);

        INSERT INTO employees_data
        (name, email, department, salary, joining_date, created_at, updated_at)
        VALUES
        (
            random_name,
            lower(random_name) || i || '@gmail.com',
            random_dept,
            random_salary,
            CURRENT_DATE - (random()*30)::INT,
            NOW(),
            NOW()
        );

    END LOOP;
END;
$$;