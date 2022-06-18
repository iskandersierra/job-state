CREATE TABLE jobs.JobState
(
    JobId CHAR(22) NOT NULL,

    Title NVARCHAR(100) NOT NULL,
    JobType NVARCHAR(100) NOT NULL,
    StepCount INT NOT NULL,
    Metadata NVARCHAR(MAX) NOT NULL,

    Progress INT NOT NULL,
    Status INT NOT NULL,
    Error NVARCHAR(MAX) NULL,

    CreatedAt DATETIME2 NOT NULL,
    UpdatedAt DATETIME2 NOT NULL,
    FinishedAt DATETIME2 NULL,

    CONSTRAINT PK__JobState PRIMARY KEY CLUSTERED (JobId ASC)
);

CREATE TABLE jobs.JobStateStep
(
    JobId CHAR(22) NOT NULL,
    StepId INT NOT NULL,

    Title NVARCHAR(100) NOT NULL,
    StepType NVARCHAR(100) NOT NULL,
    Progress INT NOT NULL,
    Status INT NOT NULL,
    Metadata NVARCHAR(MAX) NOT NULL,
    Error NVARCHAR(MAX) NULL,

    CreatedAt DATETIME2 NOT NULL,

    CONSTRAINT PK__JobStateStep PRIMARY KEY CLUSTERED (JobId ASC, StepId ASC),
);

ALTER TABLE jobs.JobStateStep
    ADD CONSTRAINT FK__JobStateStep_Job FOREIGN KEY (JobId) REFERENCES jobs.JobState(JobId);
