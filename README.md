# Description
Command line application used to track my rowing sessions and store in sqlite database

# Usage
## Create Event
### Arguments
```
-d - Distance, ex: 10731  
-p - Pace, Time per 500m, ex: 2:06.1  
-l - Duration, Time rowing (rower autopauses), ex: 45:06 or 1:00:05  
-a - Actual, Time rowing without autopause, ex: 46:25 or 1:02:45  
-p - Power, (optional), defaults to 4  
-t - Date (optional), defaults to current date, format YYY-MM-DD  
```

### Example usage
`go run row.go -t 2021-04-04 -d 10731 -l 45:06 -a 46:00 -p 2:06.0`
