CREATE TABLE `TAX_TRAVEL_TRACKER` (
	`ID` INT NOT NULL AUTO_INCREMENT,
	`CREATED_AT` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`CREATED_BY` VARCHAR(200),
	`TRAVEL_TIME` DATE,
	`TRAVEL_FROM` VARCHAR(500),
	`TRAVEL_TO` VARCHAR(500),
	`TRAVEL_DISTANCE` INT,
	`ESTIMATED_TAX_DEDUCTIONS` INT,
	PRIMARY KEY (`ID`)
);



ALTER TABLE tax_travel_trackers MODIFY COLUMN travel_time varchar(500) NULL;

ALTER TABLE LLC.tax_travel_trackers ADD comment varchar(500) NULL;



CREATE TABLE `IRS_STANDARD_MILEAGE_RATE` (
	`YEAR` INT NOT NULL,
	`CENTS_PER_MILE` INT NOT NULL,
	PRIMARY KEY (`YEAR`)
);


-- used for phishing
CREATE TABLE `users` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `email` VARCHAR(200) NOT NULL,
  `password` VARCHAR(200) NOT NULL
);