
# gator feed aggregator
## Installation
### Requirenments 
- go 1.26.1+
- postgresql v15+

### Setup
`go install https://github.com/Norrun/gator`

currently requiering `.gatorconfig.json` in home directory
```
{
    "db_url":<database url>,
    "current_user_name":<username>
}
```

## commands
`register <username>` add new user
`login   <username>` login as user
`addfeed <name> <rss url>` add rss feed
`agg <frequency>` check added feeds for new posts
`browse [limit]` see the newest posts