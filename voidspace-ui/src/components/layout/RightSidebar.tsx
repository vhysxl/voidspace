"use client";

import { Search } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

const suggestedUser = {
  id: 41,
  name: "Lotan",
  username: "vhysxl",
  avatar: "https://storage.googleapis.com/assets_voidspace/avatars/1775756391697344648.jpg",
  bio: "Developernya bang"
};

export default function RightSidebar() {
  const router = useRouter();
  const [searchQuery, setSearchQuery] = useState("");

  const handleSearch = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" && searchQuery.trim()) {
      router.push(`/explore?q=${encodeURIComponent(searchQuery.trim())}`);
    }
  };

  return (
    <aside className="hidden lg:flex flex-col w-[350px] h-screen fixed right-0 top-0 border-l border-foreground/10 bg-background z-40 p-6 space-y-8 overflow-y-auto scrollbar-hide">
      {/* Search Bar */}
      <div className="relative group">
        <div className="absolute left-4 top-1/2 -translate-y-1/2 text-foreground/40 group-focus-within:text-foreground transition-colors">
          <Search size={18} />
        </div>
        <input
          type="text"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          onKeyDown={handleSearch}
          placeholder="Search Void"
          className="w-full bg-foreground/5 border border-foreground/10 rounded-full py-3 pl-12 pr-4 text-sm tracking-[1px] outline-none focus:border-foreground/30 transition-all placeholder:text-foreground/20"
        />
      </div>

      {/* Who to follow */}
      <div className="bg-foreground/5 border border-foreground/5 rounded-sm p-6 space-y-6">
        <h2 className="font-space-grotesk font-bold text-lg tracking-[-0.5px] uppercase text-foreground">
          Who to follow
        </h2>

        <div className="space-y-5">
          <Link
            href={`/profile/${suggestedUser.username}`}
            className="flex items-center justify-between group cursor-pointer hover:bg-foreground/[0.02] p-2 -m-2 rounded-sm transition-colors"
          >
            <div className="flex items-center gap-3">
              <div className="size-10 rounded-full bg-foreground/10 border border-foreground/5 flex items-center justify-center overflow-hidden">
                {suggestedUser.avatar ? (
                  <img src={suggestedUser.avatar} alt={suggestedUser.name} className="w-full h-full object-cover" />
                ) : (
                  <span className="text-xs font-bold text-foreground/20 uppercase">
                    {suggestedUser.username.slice(0, 2)}
                  </span>
                )}
              </div>
              <div className="flex flex-col">
                <span className="text-foreground text-[13px] font-bold tracking-[0.5px] group-hover:underline">
                  {suggestedUser.name}
                </span>
                <span className="text-foreground/40 text-[11px] tracking-[1px]">
                  @{suggestedUser.username}
                </span>
              </div>
            </div>
            <button className="bg-foreground text-background px-4 py-1.5 rounded-full text-[11px] font-bold uppercase tracking-[1px] hover:opacity-90 transition-colors">
              Follow
            </button>
          </Link>
        </div>

        <button className="text-foreground/40 hover:text-foreground text-[11px] font-bold uppercase tracking-[2px] transition-colors w-full text-left pt-2">
          Show more
        </button>
      </div>

    </aside>
  );
}
