-- schema.sql is a file that contains the SQL commands to create the database schema.

-- Table for Students
CREATE TABLE Students (
                          ID CHAR(26) PRIMARY KEY,
                          Name VARCHAR(255) NOT NULL,
                          Email VARCHAR(255) NOT NULL UNIQUE,
                          Password VARCHAR(255) NOT NULL,  -- Store hashed and salted password
                          RegistryDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for Instructors
CREATE TABLE Instructors (
                             ID CHAR(26) PRIMARY KEY,
                             Name VARCHAR(255) NOT NULL,
                             Email VARCHAR(255) NOT NULL UNIQUE,
                             Password VARCHAR(255) NOT NULL,  -- Store hashed and salted password
                             RegistryDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             AreaOfExpertise VARCHAR(255)
);

-- Table for Courses
CREATE TABLE Courses (
                         ID CHAR(26) PRIMARY KEY,
                         Title VARCHAR(255) NOT NULL,
                         Description TEXT,
                         CourseDate DATE,
                         InstructorID CHAR(26),
                         FOREIGN KEY (InstructorID) REFERENCES Instructors(ID)
);

-- Table for Enrollments
CREATE TABLE Enrollments (
                            ID CHAR(26) PRIMARY KEY,
                            StudentID CHAR(26),
                            CourseID CHAR(26),
                            EnrollmentDate DATE,
                            Grade INTEGER,
                            FOREIGN KEY (StudentID) REFERENCES Students(ID) ON DELETE CASCADE,
                            FOREIGN KEY (CourseID) REFERENCES Courses(ID) ON DELETE CASCADE
);