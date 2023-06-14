import Navbar from "@/components/navbar";
import Link from "next/link";
import { MouseEvent, useEffect } from "react";

import { useAtom } from "jotai";
import { userAtom } from "@/lib/store";

import { useQuery, useMutation } from "@tanstack/react-query";
import axios from "axios";

import { VoucherData } from "@/types/type";

import Router from "next/router";

export default function Home() {
  const [user, setUser] = useAtom(userAtom);

  const voucherData = useQuery({
    queryKey: ["voucherData"],
    queryFn: async () => {
      const res = await axios.get("http://localhost:8080/voucher/list");
      // console.log(res);
      return res.data as VoucherData[];
    },
  });

  return (
    <main>
      <Navbar />
      <div className="mt-24 flex flex-col gap-2">
        {user.isAdmin ? (
          <Link href="/admin/addvoucher">
            <button className=" py-2 px-4 relative left-3/4 max-w-fit border border-black rounded shadow-md">
              Tambah Voucher
            </button>
          </Link>
        ) : (
          <div></div>
        )}
        <div className="h-1/3 mx-36 px-9 bg-slate-200 py-8 ">
          <div className="  flex flex-col gap-4 place-content-center">
            {voucherData.isSuccess ? (
              voucherData.data
                .filter((data) => data.quantity > 0)
                .map((data) => {
                  return (
                    <VoucherCard
                      user={user}
                      data={data}
                      key={data.id}
                      id={data.id}
                    />
                  );
                })
            ) : (
              <p>Tidak ada data</p>
            )}
          </div>
        </div>
      </div>
    </main>
  );
}

function VoucherCard({
  user,
  data,
  id,
}: // handlers,
{
  user: { isAdmin: boolean; isLoggedIn: boolean };
  data: VoucherData;
  id: number;
}) {
  const buyMutation = useMutation({
    mutationFn: async (id: number) => {
      const res = await axios.post(
        `http://localhost:8080/transaction/buy/${id}`,
        {},
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        }
      );
      return res;
    },
  });

  const deleteMutation = useMutation({
    mutationFn: async (id: number) => {
      return axios.delete(`http://localhost:8080/voucher/delete/${id}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });
    },
  });

  const handlers = {
    buyVoucher: (e: MouseEvent<HTMLButtonElement>) => {
      e.preventDefault();
      buyMutation.mutate(id, {
        onSuccess: () => {
          alert("Voucher berhasil dibeli");
        },
      });
    },
    deleteVoucher: (e: MouseEvent<HTMLButtonElement>) => {
      e.preventDefault();
      deleteMutation.mutate(id, {
        onSuccess: () => {
          alert("Voucher berhasil dihapus");
        },
      });
    },
  };
  return (
    <div className="bg-slate-300 rounded shadow h-20 mx-32 px-24 flex justify-between place-items-center">
      <div className="flex flex-col min-w-fit">
        <p>
          {data.name} - Rp.{data.price}
        </p>
        <p className="text-xs italic">
          {data.code} - valid till : {data.expiry_date}
        </p>
        <div className="text-xs ">Stock : {data.quantity}</div>
      </div>
      <div className="flex gap-1">
        {user.isLoggedIn && !user.isAdmin ? (
          <button
            onClick={handlers.buyVoucher}
            className="bg-blue-600 hover:bg-blue-400 transition text-white text-sm px-3 py-2 shadow rounded-md"
          >
            Beli
          </button>
        ) : (
          <></>
        )}
        {user.isAdmin ? (
          <>
            <button className="bg-yellow-600 hover:bg-yellow-500 transition text-white text-sm px-3 py-2 shadow rounded-md">
              Ubah
            </button>
            <button
              className="bg-red-600 hover:bg-red-400 transition text-white text-sm px-3 py-2 shadow rounded-md"
              onClick={handlers.deleteVoucher}
            >
              Hapus
            </button>
          </>
        ) : (
          <></>
        )}
      </div>
    </div>
  );
}
