export default function AuthFooter() {
  return (
    <div className="mt-24 text-center space-y-4 font-manrope">
      <p className="text-sm text-[#666] tracking-[0.9px] transition-colors uppercase">
        {new Date().getFullYear()} © VOIDSPACE
      </p>
    </div>
  );
}
