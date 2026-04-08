"use client";

import { Search } from "lucide-react";
import Button from "@/components/ui/Button";

const dummyUsers = [
  { id: 1, name: "Cosmic Voyager", username: "voyager_01", avatar: null },
  { id: 2, name: "Nebula Dreams", username: "nebula_x", avatar: null },
  { id: 3, name: "Star Dust", username: "stardust_pixel", avatar: null },
];

export default function RightSidebar() {
  return (
    <aside className="hidden lg:flex flex-col w-[350px] h-screen fixed right-0 top-0 border-l border-foreground/10 bg-background z-40 p-6 space-y-8 overflow-y-auto scrollbar-hide">
      {/* Search Bar */}
      <div className="relative group">
        <div className="absolute left-4 top-1/2 -translate-y-1/2 text-foreground/40 group-focus-within:text-foreground transition-colors">
          <Search size={18} />
        </div>
        <input 
          type="text"
          placeholder="SEARCH VOID"
          className="w-full bg-foreground/5 border border-foreground/10 rounded-full py-3 pl-12 pr-4 text-sm tracking-[1px] uppercase outline-none focus:border-foreground/30 transition-all placeholder:text-foreground/20"
        />
      </div>

      {/* Who to follow */}
      <div className="bg-foreground/5 border border-foreground/5 rounded-sm p-6 space-y-6">
        <h2 className="font-space-grotesk font-bold text-lg tracking-[-0.5px] uppercase text-foreground">
          Who to follow
        </h2>
        
        <div className="space-y-5">
          {dummyUsers.map((user) => (
            <div key={user.id} className="flex items-center justify-between group">
              <div className="flex items-center gap-3">
                <div className="size-10 rounded-full bg-foreground/10 border border-foreground/5 flex items-center justify-center text-xs font-bold text-foreground/20 uppercase">
                  {user.username.slice(0, 2)}
                </div>
                <div className="flex flex-col">
                  <span className="text-foreground text-[13px] font-bold tracking-[0.5px] group-hover:underline cursor-pointer">
                    {user.name}
                  </span>
                  <span className="text-foreground/40 text-[11px] uppercase tracking-[1px]">
                    @{user.username}
                  </span>
                </div>
              </div>
              <button className="bg-foreground text-background px-4 py-1.5 rounded-full text-[11px] font-bold uppercase tracking-[1px] hover:opacity-90 transition-colors">
                Follow
              </button>
            </div>
          ))}
        </div>

        <button className="text-foreground/40 hover:text-foreground text-[11px] font-bold uppercase tracking-[2px] transition-colors w-full text-left pt-2">
          Show more
        </button>
      </div>

    </aside>
  );
}
