import { useState } from "react";
import { useRouter } from "next/router";

import { useMutation } from "@tanstack/react-query";
import axios from "axios";

export default function AddVoucher() {
  const [voucherBody, setVoucherBody] = useState({
    name: "",
    code: "",
    value: 0,
    price: 0,
    quantity: 0,
    expiry_date: "",
  });

  const router = useRouter();

  interface VoucherPayload {
    name: string;
    code: string;
    value: number;
    price: number;
    quantity: number;
    expiry_date: string;
  }

  const createPost = useMutation({
    mutationFn: async (payload: VoucherPayload) => {
      const data = await axios.post(
        "http://localhost:8080/manage/create",
        payload,
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        }
      );
      return data;
    },
  });

  return (
    <div className="flex place-content-center ">
      <div className="w-full max-w-md pt-16">
        <form className="bg-slate-200 shadow rounded px-8 pt-6 pb-8 mb-4">
          <h1 className="mb-4 text-center text-xl font-bold">Add a Voucher</h1>

          <div className="flex gap-2">
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Name
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="name"
                type="text"
                placeholder="Voucher Name"
                onChange={(e) =>
                  setVoucherBody({ ...voucherBody, name: e.target.value })
                }
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Code
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="code"
                type="text"
                placeholder="Voucher Code"
                onChange={(e) =>
                  setVoucherBody({ ...voucherBody, code: e.target.value })
                }
              />
            </div>
          </div>

          <div className="flex gap-2">
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Value
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="value"
                type="number"
                placeholder="Voucher Value"
                onChange={(e) =>
                  setVoucherBody({
                    ...voucherBody,
                    value: Number(e.target.value),
                  })
                }
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Price
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="price"
                type="number"
                placeholder="Voucher Buying Price"
                onChange={(e) =>
                  setVoucherBody({
                    ...voucherBody,
                    price: Number(e.target.value),
                  })
                }
              />
            </div>
          </div>

          <div className="flex gap-2">
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Quantity
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="quantity"
                type="number"
                placeholder="Voucher Quantity"
                onChange={(e) =>
                  setVoucherBody({
                    ...voucherBody,
                    quantity: Number(e.target.value),
                  })
                }
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2">
                Expiry Date
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="expire"
                type="date"
                onChange={(e) =>
                  setVoucherBody({
                    ...voucherBody,
                    expiry_date: e.target.value,
                  })
                }
              />
            </div>
          </div>

          <div className="flex items-center justify-between">
            <button
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="button"
              onClick={(e) => {
                e.preventDefault();

                createPost.mutate(voucherBody, {
                  onSuccess: () => {
                    alert("Voucher Succesfully Created");

                    router.push("/");
                  },
                  onError: () => {
                    alert("Failed to create voucher");
                  },
                });
              }}
            >
              Create New Voucher
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
