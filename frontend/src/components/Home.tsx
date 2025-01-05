/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";

const Home = () => {
  const navigate = useNavigate();
  const [blogs, setBlogs] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get("http://127.0.0.1:3000/getBlog");
        setBlogs(response.data.data);
      } catch (error) {
        console.error("Error fetching blogs:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();

    const token = localStorage.getItem("token");
    if (!token) {
      alert("You are not logged in. Redirecting to login page.");
      navigate("/login");
    }
  }, [navigate]);

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      {loading ? (
        <p>Loading blogs...</p>
      ) : (
        blogs.map((blog: any) => (
          <Card key={blog._id} className="shadow-lg mb-4">
            <CardHeader>
              <CardTitle>{blog.title}</CardTitle>
            </CardHeader>
            <CardContent>
              <p>{blog.content}</p>
              {blog.image && <img src={blog.image} alt="Blog" />}
            </CardContent>
          </Card>
        ))
      )}
      <div className="max-w-2xl mx-auto mt-8">
        <Button onClick={() => navigate('/blog')} className="w-full">
          Create New Blog Post
        </Button>
      </div>
    </div>
  );
};

export default Home;
