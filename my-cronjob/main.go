/*
*************************************************************************

I want it have a little workload, just focus the program design.
- unit is hour
- don't care time_zone
- don't set the time pausing between jobs in same schedule (in real, the
	field will make sure process more stable, I think :v)
- job only pause by people,
	in real, it can pause by other job or schedule.

Quick requirements (Job):
- Can set a job (not started, just set the job's params): 2 scenarios
  - 1: specific time, time = 0 -> run immediately
  - 2: run by time interval, interval > 0 (case loop ignored)
  - Job's status: NOT_STARTED, LOCKING, RUNNING, COMPLETED, IGNORE

- 1 schedule = n jobs
  - Can pause, continue, delete 1 job.
  When continue but the current_time > run_at -> job's status = "IGNORE"
  - Can start, clear 1 schedule: start schedule mean starting first job

Technical specification:
- Job
- Schedule

*************************************************************************
*/

package main

import "time"

// hack the maximum if needed
const MAX_JOB = 9999

type Job struct {
	err      error
	run_at   time.Duration
	last_run time.Time
	next_run time.Time
	lock     bool
	interval uint64
}

type Schedule struct {
}

// Create new job with time interval
func NewJob(interval uint64) *Job {
	return &Job{
		interval: interval,
		last_run: time.Unix(0, 0),
		next_run: time.Unix(0, 0),
	}
}
