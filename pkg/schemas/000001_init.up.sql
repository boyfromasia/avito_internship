CREATE TABLE Users
(
    UserId  int     NOT NULL        PRIMARY KEY,
    Balance float   NOT NULL,
    Reserve float   NOT NULL
);


CREATE TABLE Purchases
(
    PurchaseId int       NOT NULL   PRIMARY KEY,
    Name   varchar(255)  NOT NULL
);


CREATE TABLE Orders
(
    OrderId         SERIAL         PRIMARY KEY,
    UserId          int            REFERENCES Users     (UserId)      NOT NULL,
    PurchaseId      int            REFERENCES Purchases (PurchaseId),
    Price           float          NOT NULL,
    Comment         varchar(255)   NOT NULL,
    TimeCreated     timestamp      NOT NULL,
    StatusOrder     varchar(255)   NOT NULL
);