import { BrowserRouter, Route, Routes } from 'react-router-dom';
import SignupPage from './components/SignUp';
import OtpPage from './components/Otp';
import LoginPage from './components/Login';
import Home from './components/Home';
import NotFound from './components/NotFound';


export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/signup" element={<SignupPage />} />
        <Route path="/otp" element={<OtpPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/home" element={<Home />} />
        <Route path="*" element={<NotFound />} />       
        {/* Add more routes as needed */}
      </Routes>
    </BrowserRouter>
  );
}