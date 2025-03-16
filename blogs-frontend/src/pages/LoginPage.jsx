import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { TextField, Button, Container, Typography } from "@mui/material";
import axios from "axios";

const LoginPage = () => {
  const [formData, setFormData] = useState({ email: "", password: "" });
  const [error, setError] = useState("");
  const navigate = useNavigate();

  console.log("Rendered LoginPage"); // This will help track re-renders

  // Handles input change
  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  // Handles form submission
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    try {
      const response = await axios.post("http://localhost:5000/api/login", formData);
      localStorage.setItem("token", response.data.token);
      navigate("/");
    } catch (err) {
      setError("Invalid email or password");
    }
  };

  return (
    <Container maxWidth="ms" className="flex flex-col items-center justify-center h-screen ">
      <Typography variant="h4" className="mr-4">Login</Typography>
      
      {error && <Typography color="error">{error}</Typography>}

      <form onSubmit={handleSubmit} className="w-full">
        <TextField
          label="Email"
          name="email"
          variant="outlined"
          fullWidth
          className="mb-4"
          value={formData.email}
          onChange={handleChange}
        />

        <TextField
          label="Password"
          name="password"
          type="password"
          variant="outlined"
          fullWidth
          className="mb-4"
          value={formData.password}
          onChange={handleChange}
        />

        <Button type="submit" variant="contained" color="primary" fullWidth>
          Login
        </Button>
      </form>
    <Button onClick={()=>navigate("/register")} 
    variant="outlined" 
    color="primary" 
    className="mt-6"
    fullWidth > Create Account </Button>
    </Container>
  );
};

export default LoginPage;
