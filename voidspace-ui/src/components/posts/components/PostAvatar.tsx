import Link from "next/link";

interface PostAvatarProps {
  username: string;
  avatarUrl?: string;
}

export default function PostAvatar({ username, avatarUrl }: PostAvatarProps) {
  return (
    <Link
      href={`/profile/${username}`}
      className="shrink-0 pt-1"
      onClick={(e) => e.stopPropagation()}
    >
      <div className="size-10 md:size-12 rounded-full bg-void-gray border border-foreground/5 overflow-hidden flex items-center justify-center">
        {avatarUrl ? (
          <img src={avatarUrl} alt="" className="w-full h-full object-cover" />
        ) : (
          <span className="text-xs font-bold uppercase text-foreground/20">
            {username.slice(0, 2)}
          </span>
        )}
      </div>
    </Link>
  );
}
