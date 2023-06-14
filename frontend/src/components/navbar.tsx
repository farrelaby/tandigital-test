import Link from "next/link";

import { useAtom } from "jotai";
import { userAtom } from "@/lib/store";

import Router from "next/router";
import { useEffect } from "react";
import jwtDecode from "jwt-decode";

type decodedToken = {
  role: string;
  id?: string;
};

export default function Navbar() {
  const [user, setUser] = useAtom(userAtom);

  useEffect(() => {
    const jwToken = String(localStorage.getItem("token"));

    if (jwToken != "empty") {
      const decodedToken: decodedToken = jwtDecode(jwToken);
      if (decodedToken.role === "admin") {
        setUser({ ...user, isLoggedIn: true, isAdmin: true });
      } else {
        setUser({
          isLoggedIn: true,
          isAdmin: false,
          id: decodedToken.id as string,
        });
      }
    }
  }, []);

  const handleLogout = () => {
    localStorage.setItem("token", "empty");
    setUser({ isLoggedIn: false, isAdmin: false, id: "" });
    Router.reload();
  };

  return (
    <nav className="h-20 w-full px-44 pt-1 mb-8 bg-[#b0dfcd]  flex flex-row justify-between place-items-center ">
      <Link href="/">
        <div className="font-bold">Voucher Marketplace</div>
      </Link>
      {user.isLoggedIn && !user.isAdmin ? (
        <Link href="/myvoucher">
          <div className="border border-black rounded shadow-md py-2 px-4">
            My Vouchers
          </div>
        </Link>
      ) : (
        <></>
      )}
      <ul className="flex flex-row gap-11 place-items-center">
        <div className="flex flex-row gap-5">
          {user.isLoggedIn ? (
            <li>
              <Link href="/">
                <button
                  className="ml-2 bg-transparent font-semibold py-2 px-5 border border-black rounded-md shadow-md"
                  onClick={handleLogout}
                >
                  Log Out
                </button>
              </Link>
            </li>
          ) : (
            <>
              <li>
                <Link href="/login">
                  <button className="ml-2 bg-transparent font-semibold py-2 px-5 border border-black rounded-md shadow-md">
                    Log In
                  </button>
                </Link>
              </li>
              <li>
                <Link href="/register">
                  <button className="bg-black text-white font-semibold py-2 px-5 border border-black rounded-md shadow-md">
                    Sign Up
                  </button>
                </Link>
              </li>
            </>
          )}
        </div>
      </ul>
    </nav>
  );
}
