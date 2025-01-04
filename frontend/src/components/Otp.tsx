import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function OtpPage() {
  const [otp, setOtp] = useState('');
  const navigate = useNavigate();

  const handleVerifyOtp = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!otp) {
      alert('Please enter OTP!');
      return;
    }

    const otpData = { val: parseInt(otp, 10) }; // Match the expected `val` field in your backend

    try {
      const response = await fetch('http://127.0.0.1:3000/checkotp', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(otpData),
      });

      if (response.ok) {
        const result = await response.json();
        alert(result.message); // Notify the user
        navigate('/login'); // Redirect to login on success
      } else {
        const error = await response.json();
        alert(error.message); // Notify the user of the error
        navigate('/signup'); // Redirect back to signup on failure
      }
    } catch (err) {
      console.error('Error:', err);
      alert('Something went wrong, please try again!');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
        <h2 className="text-2xl font-semibold text-center mb-6">OTP Verification</h2>
        <form onSubmit={handleVerifyOtp}>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2">Enter OTP</label>
            <input
              type="text"
              value={otp}
              onChange={(e) => setOtp(e.target.value)}
              className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-500 text-white font-medium py-2 px-4 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400"
          >
            Verify OTP
          </button>
        </form>
      </div>
    </div>
  );
}

export default OtpPage;
