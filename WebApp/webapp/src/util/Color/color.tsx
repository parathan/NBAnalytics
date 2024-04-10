export function perc2color(perc: number) {
    let min: number = 0;
    let max: number = 100;
    var base = (max - min);

    if (perc > 100) {
        return '#00ff00';
    } else if (perc < 0) {
        return '#ff0000';
    }

    if (base == 0) { perc = 100; }
    else {
        perc = (perc - min) / base * 100; 
    }
    var r, g, b = 0;
    if (perc < 50) {
        r = 255;
        g = Math.round(5.1 * perc);
    }
    else {
        g = 255;
        r = Math.round(510 - 5.10 * perc);
    }
    var h = r * 0x10000 + g * 0x100 + b * 0x1;
    return '#' + ('000000' + h.toString(16)).slice(-6);
}