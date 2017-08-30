# email-server

This is a simple server written in Go. It's function is to receive http 
requests generated from [IndieExpo][IndieExpo], [AI-Wins][aiwins], 
and [my personal portfolio site][portfolio], and forward them as emails
using gmail's smtp server. In addition to Go, it uses the mux package from
the Gorilla web toolkit to process and route http requests. It is currently
sitting behind a Caddy web server on a tiny Ubuntu compute instance on the
Google Cloud Platform. It is prefectly happy there. 

[IndieExpo]: https://github.com/TheRoyalTnetennba/indie-expo
[aiwins]: https://github.com/TheRoyalTnetennba/ai-wins
[portfolio]: https://github.com/TheRoyalTnetennba/portfolio