ALTER TABLE jobs.JobStateStep
    DROP CONSTRAINT FK__JobStateStep_Job;

DROP TABLE jobs.JobStateStep;

DROP TABLE jobs.JobState;
