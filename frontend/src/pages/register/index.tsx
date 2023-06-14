import { useState } from "react";
import { useRouter } from "next/router";

import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { useAtom } from "jotai";
import { userAtom } from "@/lib/store";

export default function Signup() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const router = useRouter();

  type Token = {
    data: {
      token: string;
    };
  };

  const loginPost = useMutation({
    mutationFn: async (payload: { email: string; password: string }) => {
      const data: Token = await axios.post(
        "http://localhost:8080/consumer/register",
        payload
      );
      return data;
    },
  });

  return (
    <div className="flex place-content-center ">
      <div className="w-full max-w-xs pt-16">
        <form className="bg-slate-200 shadow rounded px-8 pt-6 pb-8 mb-4">
          <h1 className="mb-4 text-center text-xl font-bold">
            Register an Account
          </h1>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Email
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="username"
              type="text"
              placeholder="Email"
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <div className="mb-6">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Password
            </label>
            <input
              className="shadow appearance-none  rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
              id="password"
              type="password"
              placeholder="******************"
              onChange={(e) => setPassword(e.target.value)}
            />
            <p className="text-red-500 text-xs italic">
              Please insert your password
            </p>
          </div>
          <div className="flex items-center justify-between">
            <button
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="button"
              onClick={(e) => {
                e.preventDefault();
                loginPost.mutate(
                  { email, password },
                  {
                    onSuccess: () => {
                      router.push("/login");
                    },
                  }
                );
              }}
            >
              Register
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
