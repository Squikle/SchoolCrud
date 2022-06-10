CREATE TABLE Schools
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Address VARCHAR(50) UNIQUE NOT NULL,
    CoursePrice DECIMAL(8,2) NOT NULL
);

CREATE TABLE Persons
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    FistName VARCHAR(30) NOT NULL,
    LastName VARCHAR(30) NOT NULL,
    Phone VARCHAR(20) UNIQUE NOT NULL,
    Email VARCHAR(50) UNIQUE NULL,
    Sex VARCHAR(20) NULL
);

CREATE TABLE Instructors
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    PersonId INT UNIQUE NOT NULL,
    SchoolId INT NULL,
    CONSTRAINT instructors_persons_fk
        FOREIGN KEY (PersonId) REFERENCES Persons (Id) ON DELETE CASCADE,
    CONSTRAINT instructors_schools_fk
        FOREIGN KEY (SchoolId) REFERENCES Schools (Id) ON DELETE SET NULL
);

CREATE TABLE Students
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    PersonId INT UNIQUE NOT NULL,
    SchoolId INT NULL,
    CONSTRAINT students_persons_fk
        FOREIGN KEY (PersonId) REFERENCES Persons (Id) ON DELETE CASCADE,
    CONSTRAINT students_schools_fk
        FOREIGN KEY (SchoolId) REFERENCES Schools (Id) ON DELETE SET NULL
);

CREATE TABLE `Groups`
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    InstructorId INT NOT NULL,
    SchoolId INT NOT NULL,
    CONSTRAINT groups_instructors_fk
        FOREIGN KEY (InstructorId) REFERENCES Instructors (Id) ON DELETE NO ACTION,
    CONSTRAINT groups_schools_fk
        FOREIGN KEY (SchoolId) REFERENCES Schools (Id) ON DELETE NO ACTION
);

CREATE TABLE StudentsGroups
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    StudentId INT NOT NULL,
    GroupId INT NOT NULL,
    CONSTRAINT studentsGroups_students_fk
        FOREIGN KEY (StudentId) REFERENCES Students (Id) ON DELETE CASCADE,
    CONSTRAINT studentsGroups_groups_fk
        FOREIGN KEY (GroupId) REFERENCES `Groups` (Id) ON DELETE CASCADE
);

CREATE TABLE Cars
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Brand VARCHAR(50) NOT NULL,
    Model VARCHAR(50) NOT NULL,
    Number VARCHAR(10) NOT NULL,
    Year INT NOT NULL,
    Color VARCHAR(20) NOT NULL
);

CREATE TABLE Payments
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    StudentId INT NOT NULL,
    SchoolId INT NOT NULL,
    Amount DECIMAL(8,2) NOT NULL,
    CONSTRAINT payments_students_fk
        FOREIGN KEY (StudentId) REFERENCES Students (Id) ON DELETE NO ACTION,
    CONSTRAINT payments_schools_fk
        FOREIGN KEY (SchoolId) REFERENCES Schools (Id) ON DELETE NO ACTION
);

CREATE TABLE Topics
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50) UNIQUE NOT NULL,
    Description VARCHAR(50) NULL
);

CREATE TABLE Lessons
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    InstructorId INT NOT NULL,
    TopicId INT NOT NULL,
    CONSTRAINT lessons_instructors_fk
        FOREIGN KEY (InstructorId) REFERENCES Instructors (Id) ON DELETE NO ACTION,
    CONSTRAINT lessons_topics_fk
        FOREIGN KEY (TopicId) REFERENCES Topics (Id) ON DELETE NO ACTION
);

CREATE TABLE PracticeLessons
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    LessonId INT UNIQUE NOT NULL,
    StudentId INT NOT NULL,
    CarId INT NOT NULL,
    CONSTRAINT practiceLessons_lessons_fk
        FOREIGN KEY (LessonId) REFERENCES Lessons (Id) ON DELETE CASCADE,
    CONSTRAINT practiceLessons_students_fk
        FOREIGN KEY (StudentId) REFERENCES Students (Id) ON DELETE NO ACTION,
    CONSTRAINT practiceLessons_cars_fk
        FOREIGN KEY (CarId) REFERENCES Cars (Id) ON DELETE NO ACTION
);

CREATE TABLE TheoryLessons
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    LessonId INT UNIQUE NOT NULL,
    GroupId INT NOT NULL,
    CONSTRAINT theoryLessons_lessons_fk
        FOREIGN KEY (LessonId) REFERENCES Lessons (Id) ON DELETE CASCADE,
    CONSTRAINT theoryLessons_groups_fk
        FOREIGN KEY (GroupId) REFERENCES `Groups` (Id) ON DELETE NO ACTION
);

INSERT INTO Cars(Brand, Model, Number, Year, Color)
VALUES ('Skoda', 'Fabia', 'AK9565AK', '2007', 'White'),
       ('Chevrolet', 'Aveo', 'AM4212AD', '2009', 'Blue');

INSERT INTO Schools(Address, CoursePrice)
VALUES ('General Arington, Virginia 22204', 11500),
       ('Pershing Drive, Virginia 22261', 14000);

INSERT INTO Persons(FistName, LastName, Phone, Email, Sex)
VALUES ('Max', 'Fedorenko', '+380952495125', 'MaxLo@email.com', null),
       ('Pavlo', 'Skydan', '+380956816153', 'PavlO@email.com', 'Male'),
       ('Dovhalov', 'Mykhailo', '+380956819853', 'm.Dovhalov@email.com', 'Male'),
       ('Lena', 'Golovach', '+380956819123', 'Lenich@ukrmail.com', 'Female'),
       ('Anton', 'Davidovich', '+380956819532', 'Anton.Scr@gmail.com', 'Female');

INSERT INTO Instructors(PersonId, SchoolId)
VALUES (3, 1),
       (4, 2);

INSERT INTO Students(PersonId, SchoolId)
VALUES (1, 1),
       (2, 2),
       (5, 2);

INSERT INTO `Groups`(InstructorId, SchoolId)
VALUES (1, 1),
       (2, 2);

INSERT INTO StudentsGroups(StudentId, GroupId)
VALUES (1, 1),
       (2, 2),
       (3, 2);

INSERT INTO Topics(Name, Description)
VALUES ('Reversal moving', null),
       ('Parking', 'Student learning to park a car'),
       ('SDA. Signs', null),
       ('SDA. Lightning', null);

INSERT INTO Lessons(InstructorId, TopicId)
VALUES (1, 3),
       (1, 4),
       (1, 1),
       (1, 2),
       (2, 3),
       (2, 4),
       (2, 1);

INSERT INTO TheoryLessons(LessonId, GroupId)
VALUES (1, 1),
       (2, 1),
       (5, 2),
       (6, 2);

INSERT INTO PracticeLessons(LessonId, StudentId, CarId)
VALUES (3, 1, 1),
       (4, 2, 1),
       (7, 3, 2);

INSERT INTO Payments(StudentId, SchoolId, Amount)
VALUES (1, 1, 11500),
       (2, 1, 7000),
       (2, 1, 7000),
       (3, 1, 14000);