-- Create task table
CREATE TABLE IF NOT EXISTS task (
    task_id MEDIUMINT AUTO_INCREMENT,
    name varchar(50) NOT NULL,
    PRIMARY KEY (task_id)
);