<template>
  <div class="min-h-screen bg-black flex items-center justify-center p-4 relative overflow-hidden font-inter selection:bg-white/20 selection:text-white">
    <!-- Aesthetic Decorative Elements (Figma-accurate) -->
    
    <!-- Top & Bottom Gradient Lines -->
    <div class="absolute top-0 left-0 right-0 h-px bg-gradient-to-r from-transparent via-white/20 to-transparent z-20"></div>
    <div class="absolute bottom-0 left-0 right-0 h-px bg-gradient-to-r from-transparent via-white/20 to-transparent z-20"></div>

    <!-- Diagonal Overlays -->
    <div class="absolute left-[-60px] top-1/2 -translate-y-1/2 rotate-[-45deg] opacity-40 hidden xl:block">
      <div class="bg-white/5 h-1 w-[200px]"></div>
    </div>
    <div class="absolute right-[-60px] top-1/2 -translate-y-1/2 rotate-[-45deg] opacity-40 hidden xl:block">
      <div class="bg-white/5 h-1 w-[200px]"></div>
    </div>

    <!-- Background Glow -->
    <div class="absolute inset-0 pointer-events-none">
      <div class="absolute inset-0 opacity-[0.03] bg-[radial-gradient(circle_at_center,_rgba(255,255,255,0.1)_0%,_rgba(255,255,255,0)_70%)]"></div>
    </div>

    <!-- Main Content -->
    <div class="relative z-10 w-full max-w-[480px]">
      <slot />
    </div>

    <!-- Footer Aesthetic -->
    <div class="absolute bottom-8 left-0 right-0 flex justify-between px-8 text-[10px] font-bold tracking-[2px] text-white/20 uppercase pointer-events-none">
      <div class="flex gap-6 pointer-events-auto">
        <UIcon name="i-heroicons-globe-alt" class="w-4 h-4" />
        <UIcon name="i-heroicons-command-line" class="w-4 h-4" />
      </div>
      <div>SYSTEM VERSION 0.4.2</div>
    </div>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;300;400;500;700;900&display=swap');

.font-inter {
  font-family: 'Inter', sans-serif;
}
</style>
