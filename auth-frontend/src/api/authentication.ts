
export const signUpUser = (data: any) => {
    const url = "http://localhost:8089/user/sign-up"
    const response = fetch(url, {
        method: "POST",
        body: data
    })
}