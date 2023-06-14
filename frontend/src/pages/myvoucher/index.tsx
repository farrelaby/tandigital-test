import Navbar from "@/components/navbar";

import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { VoucherData } from "@/types/type";

export default function MyVoucher() {
  const myVoucherData = useQuery({
    queryKey: ["myVoucherData"],
    queryFn: async () => {
      const res = await axios.get("http://localhost:8080/transaction/list", {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });
      return res.data as VoucherData[];
    },
  });

  return (
    <div>
      <Navbar />
      <div className="h-1/3 mx-36 px-9 bg-slate-200 py-8">
        <div className="  flex flex-col gap-4 place-content-center">
          {myVoucherData.isSuccess ? (
            myVoucherData.data.map((data) => {
              return <VoucherCard data={data} key={data.id} />;
            })
          ) : (
            <p>Tidak ada data</p>
          )}
        </div>
      </div>
    </div>
  );
}

function VoucherCard({
  data,
}: // handlers,
{
  data: VoucherData;
}) {
  return (
    <div className="bg-slate-300 rounded shadow h-20 mx-32 px-24 flex justify-between place-items-center">
      <div className="flex flex-col min-w-fit">
        <p>
          {data.name} - Rp.{data.price}
        </p>
        <p className="text-xs italic">
          {data.code} - valid till : {data.expiry_date}
        </p>
      </div>
      {/* <div className="flex gap-1">
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
      </div> */}
    </div>
  );
}
