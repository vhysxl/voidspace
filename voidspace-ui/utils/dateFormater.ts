export function formatPostDate(dateStr: string) {
  const postDate = new Date(dateStr);
  const now = new Date();
  const diffMs = now.getTime() - postDate.getTime();
  const diffSeconds = Math.floor(diffMs / 1000);
  const diffMinutes = Math.floor(diffSeconds / 60);
  const diffHours = Math.floor(diffMinutes / 60);

  if (diffHours < 24) {
    if (diffHours >= 1) return `${diffHours}h ago`;
    if (diffMinutes >= 1) return `${diffMinutes}m ago`;
    return `${diffSeconds}s ago`;
  }

  // lebih dari 24 jam → tampil tanggal
  return postDate.toLocaleDateString("us-EN", {
    day: "numeric",
    month: "long",
    year: "numeric",
  });
}
