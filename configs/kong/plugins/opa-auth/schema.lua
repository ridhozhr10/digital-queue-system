local typedefs = require "kong.db.schema.typedefs"

return {
    name = "opa-auth",
    fields = {{
        protocols = typedefs.protocols_http
    }, {
        config = {
            type = "record",
            fields = {{
                url = {
                    type = "string",
                    required = true
                }
            }, {
                role_header_name = {
                    type = "string",
                    required = true,
                    default = "X-Claims-Role"
                }
            }, {
                sub_header_name = {
                    type = "string",
                    required = true,
                    default = "X-Claims-Sub"
                }
            }}
        }
    }}
}

