import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      // Redirect to login if no token
      alert("You are not logged in. Redirecting to login page.");
      navigate("/login");
    }
  }, [navigate]);

  return (
    <div>
      <h1>Welcome to the Home Page</h1>
      <p>You are logged in.</p>
    </div>
  );
};

export default Home;
