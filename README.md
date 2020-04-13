# router
路由
==

<pre>
profile, err := log.NewProfile(map[string]string{
    log.ProfileDirectory: "./log",
    log.ProfileChannel:   "router",
})

if err != nil {
    fmt.Println(err)
    return
}

stream, err := log.NewStream(profile)
if err != nil {
    fmt.Println(err)
    return
}

logger, err := log.NewLogger(stream, "test", 1)
if err != nil {
    fmt.Println(err)
    return
}

route, err := route_sample.NewRoute()
if err != nil {
    fmt.Println(err)
    return
}

handler, err := router.NewHandler(route, logger)
if err != nil {
    fmt.Println(err)
    return
}

err = http.ListenAndServe(":8080", handler)
if err != nil {
    fmt.Println(err)
    return
}

// POST http://127.0.0.1:8080/module/controller/test
</pre>
