<div align="center">
<h2>tz-gohub-readme<h2>
<p>Update github readme with parsers</p>
</div>

## Why?
Inspired by [tw93/tw93](https://github.com/tw93/tw93) and [matchai/waka-box](https://github.com/matchai/waka-box), I'd like to create a plugin system to auto-update the github profile `readme.md` by using kinds of `parser` for different data sources like `rss`, `wakatime` etc.

## Parser

A parser can parse resourse pointed by `source`. `source` could be local file path, or a remote API endpoint.
I also implement an axios-like http client for retrieving from remote RESTful APIs. For the sake of simplicity, I only implement GET and POST. Why build a wheel? Well, I'm learning Golang by building wheels.

```
type Parser interface{
// transform input string
Parse(source string) ([]byte,error) 
}
```

### TBD
- [x] parser for wakatime stat
- [ ] parser for blog rss
- [ ] Github workflow

### Why golang?
I'm learning Golang these days(I wonder that it may be easier to implement using Python or JavaScript). 

<!-- rss-start -->
<!-- rss-end -->

<!-- wakatime-start -->
<!-- wakatime-end -->