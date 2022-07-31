
export const signUpUser = (data: any) : Promise<Response> => {
    const url = "http://localhost:8089/user/sign-up"
    return fetch(url, {
        method: "POST",
        body: JSON.stringify(data)
    })
}