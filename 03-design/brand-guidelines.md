---
name: brand-guidelines
description: Applies Anthropic's official brand colors and typography to artifacts that benefit from a consistent look-and-feel.
license: Complete terms in LICENSE.txt
---

# Anthropic Brand Styling

## Overview

Use this guide to apply Anthropic's official brand identity and style resources to UI and documentation artifacts.

Keywords: branding, corporate identity, visual identity, styling, brand colors, typography, Anthropic brand, visual formatting, visual design

## Brand Guidelines

### Colors

Main Colors:
- Dark: `#141413` - Primary text and dark backgrounds
- Light: `#faf9f5` - Light backgrounds and text on dark
- Mid Gray: `#b0aea5` - Secondary elements
- Light Gray: `#e8e6dc` - Subtle backgrounds

Accent Colors:
- Orange: `#d97757` - Primary accent
- Blue: `#6a9bcc` - Secondary accent
- Green: `#788c5d` - Tertiary accent

### Typography

- Headings: Poppins (fallback: Arial)
- Body Text: Lora (fallback: Georgia)
- Note: Fonts should be pre-installed in your environment for best results

## Features

### Smart Font Application

- Apply Poppins to headings (24pt and larger)
- Apply Lora to body text
- Use fallbacks if custom fonts are unavailable
- Preserve readability across all systems

### Text Styling

- Headings (24pt+): Poppins font
- Body text: Lora font
- Smart color selection based on background
- Preserve text hierarchy and formatting

### Shape and Accent Colors

- Non-text shapes use accent colors
- Cycle through orange, blue, and green accents
- Maintain visual interest while staying on-brand

## Technical Details

### Font Management

- Use system-installed Poppins and Lora fonts when available
- Provide automatic fallback to Arial (headings) and Georgia (body)
- No font installation required - works with existing system fonts
- For best results, pre-install Poppins and Lora fonts in your environment

### Color Application

- Use RGB/HEX values for precise brand matching
- Apply colors via CSS variables or design tokens
- Maintain color fidelity across systems
