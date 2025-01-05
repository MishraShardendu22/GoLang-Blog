import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { Pencil } from "lucide-react";

const Blog = () => {
  const navigate = useNavigate();
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [imageUrl, setImageUrl] = useState('');

  const postBlog = async () => {
    const id = localStorage.getItem("user-data-id");
    const username = localStorage.getItem("user-data-username");

    if (!id || !username) {
      alert("User data not found. Please log in again.");
      return;
    }

    const formData = new FormData();
    formData.append("username", username);
    formData.append("title", title);
    formData.append("content", content);
    formData.append("imageUrl", imageUrl);
    formData.append("userid", id);

    try {
      const res = await axios.post("http://127.0.0.1:3000/makeBlog", formData);

      if (res.status == 200) {
        alert('Blog posted successfully!');
        setTitle('');
        setContent('');
        setImageUrl('');
        navigate('/home');
      } else {
        alert('Failed to post blog.');
      }
    } catch (error) {
      console.error("Error posting blog:", error);
      alert('An error occurred while posting the blog.');
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <Card className="shadow-lg">
        <CardHeader>
          <CardTitle className="text-2xl font-bold flex items-center gap-2">
            <Pencil className="w-6 h-6" />
            Create New Blog Post
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="space-y-2">
            <label className="text-sm font-medium text-gray-700">Blog Title</label>
            <Input
              type="text"
              placeholder="Enter your blog title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="w-full"
            />
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium text-gray-700">Blog Content</label>
            <Textarea
              placeholder="Write your blog content here..."
              value={content}
              onChange={(e) => setContent(e.target.value)}
              className="min-h-[200px] w-full"
            />
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium text-gray-700">Image URL</label>
            <Input
              type="text"
              placeholder="Enter image URL"
              value={imageUrl}
              onChange={(e) => setImageUrl(e.target.value)}
              className="w-full"
            />
          </div>

          <Button onClick={postBlog} className="w-full">
            Publish Blog Post
          </Button>
        </CardContent>
      </Card>
    </div>
  );
};

export default Blog;
