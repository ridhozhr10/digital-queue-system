local http = require "resty.http"
local cjson = require "cjson.safe"

local OpaAuthHandler = {
    PRIORITY = 900, -- run after jwt plugin (1000), before proxying
    VERSION = "1.0"
}

function OpaAuthHandler:access(conf)
    local path = kong.request.get_path()
    local method = kong.request.get_method()
    local role = kong.request.get_header(conf.role_header_name)
    local sub = kong.request.get_header(conf.sub_header_name)

    local body = cjson.encode({
        path = path,
        method = method,
        subject = sub,
        role = role
    })

    local httpc = http.new()
    local res, err = httpc:request_uri(conf.url, {
        method = "POST",
        body = body,
        headers = {
            ["Content-Type"] = "application/json"
        }
    })

    if not res then
        kong.log.err("OPA service call failed: ", err)
        return kong.response.exit(500, {
            message = "Authorization service unavailable"
        })
    end

    local resp = cjson.decode(res.body)
    if not resp or not resp.allow then
        return kong.response.exit(403, {
            message = "Forbidden"
        })
    end
end

return OpaAuthHandler
