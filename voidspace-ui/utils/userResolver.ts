export const resolveAvatar = (
  avatarUrl: string | null | undefined,
  username: string
) => {
  if (avatarUrl && avatarUrl.trim() !== "") {
    return avatarUrl;
  }
  return `https://ui-avatars.com/api/?name=${encodeURIComponent(
    username
  )}&background=random&size=200`;
};

export const resolveBanner = (bannerUrl: string | null | undefined) => {
  return bannerUrl && bannerUrl.trim() !== ""
    ? bannerUrl
    : "https://picsum.photos/1280/720";
};
