import { Button, Container } from "@mui/material";
import React from "react";
const createAcc = ()=>{
    return (
        <Container>
            <h3>Create Account</h3>
            <form>
                <div className="flex flex-col items-center justify-center h-screen">
                    <input type="text" placeholder="Full Name" />
                    <input type="email" placeholder="Email"/>
                    <input type="password" placeholder="Password"/>
                    <Button type="submit" variant="contained" color="primary fullwidth" fullWidth>Create Account</Button>
                </div>
            </form>
        </Container>

    );
}
export default createAcc