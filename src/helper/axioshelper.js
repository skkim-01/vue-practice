//const axios = require('axios');
import axios from "axios"

/*
### Examples
const method = "POST"
const url = "http://13.209.153.211:9999/req/ticket/approve"
const hdrs = {
    "Content-Type": "application/text"
}
const body = {
    "value1": "123",
    "value2": 123,
}
#### hdrs, body could be null
const result = await AxiosHelper.send(method, url, hdrs, body)

### when success
if (result) {
    // console.log(result)
}

### when fail
if (!result) {
    // console.log(result) // result must be ""
}


### TODO
- How to print error message?
*/

export default class AxiosHelper {
    static async send(method, url, hdrs, body) {
        let eMsg = ""
        let data = ""
        
        await axios({
            method: method, // in
            url: url,       // in
            headers: hdrs,  // in_opt
            data: body      // in_opt
        })
        .then((r) => {
            data = r.data
        })
        .catch((e) => {
            if (e.response) {
                eMsg = e.response.data
            }
            else if (e.request) {
                eMsg = e.request
            }
            else {                
                eMsg = e.message
            }
        })

        if ( !eMsg ) {
            return data
        } else {
            console.log(`Caught Exceptions with AxiosHelper:
            {
                "method": ${method}
                "url": ${url}
                "header": ${JSON.stringify(hdrs)}
                "body": ${JSON.stringify(body)}
                "error-msg": ${eMsg}
            }`)
                        throw eMsg;
        }
    }
}