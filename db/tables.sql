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


--emlpyee by salary 
SELECT *FROM employees_data WHERE salary > 50000;


-- by depatment 
SELECT *FROM employees_data WHERE department = "marketing" 


-- counting
SELECT department, COUNT(*) AS employee_count FROM employees_data GROUP by department;

--by joining date
SELECT *FROM employees_data WHERE joining_date>=NOW() - INTERVAL'30 days'

--top 5 salaray
SELECT *FROM employees_data ORDER By salary DESC LIMIT -5