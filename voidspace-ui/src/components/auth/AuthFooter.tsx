export default function AuthFooter() {
  return (
    <div className="mt-24 text-center opacity-30 space-y-4 font-manrope">
      <p className="text-[9px] text-[#666] tracking-[0.9px] uppercase">
        © 2026 CELESTIAL VOID
      </p>
      <div className="flex justify-center gap-6 text-[9px] text-[#666] tracking-[0.9px] uppercase">
        <button className="hover:text-white transition-colors">Privacy</button>
        <button className="hover:text-white transition-colors">Terms</button>
        <button className="hover:text-white transition-colors">Orbit Status</button>
      </div>
    </div>
  );
}
